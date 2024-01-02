package services

import (
	"Demo/initialize"
	"Demo/utils"
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

type loginRequest struct {
	Name     string `json:"Name"`
	PassWord string `json:"PassWord"`
}

func Login(c *gin.Context) {
	var currentRequest loginRequest
	if err := c.ShouldBindJSON(&currentRequest); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	adminName := currentRequest.Name
	if adminName != initialize.Configuration.AdminName {
		c.JSON(http.StatusBadRequest, "Wrong Admin Name")
		return
	}

	adminPassWord := currentRequest.PassWord
	if adminPassWord != initialize.Configuration.AdminPassword {
		c.JSON(http.StatusBadRequest, "Wrong Admin PassWord")
		return
	}

	c.JSON(http.StatusOK, "true")
	return
}

type typeRequest struct {
	Type string `json:"Type"`
}

func CountByType(c *gin.Context) {
	var currentRequest typeRequest
	if err := c.ShouldBindJSON(&currentRequest); err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println("Error while parsing request: ", err.Error())
		return
	}

	currentType := currentRequest.Type
	if currentCount, ok := initialize.Configuration.CountByType[currentType]; ok {
		c.JSON(http.StatusOK, currentCount)
		return
	} else {
		c.JSON(http.StatusBadRequest, "Invalid Key")
		fmt.Println("Invalid Key")
		return
	}
}

func CountByGrade(c *gin.Context) {
	var currentRequest typeRequest
	if err := c.ShouldBindJSON(&currentRequest); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	currentType := currentRequest.Type
	if currentCount, ok := initialize.Configuration.CountByGrade[currentType]; ok {
		c.JSON(http.StatusOK, currentCount)
		return
	} else {
		c.JSON(http.StatusBadRequest, "Invalid Key")
		return
	}
}

type countDateRequest struct {
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
}

func CountDate(c *gin.Context) {
	var currentRequest countDateRequest
	if err := c.BindJSON(&currentRequest); err != nil {
		c.JSON(http.StatusBadRequest, "Error parsing JSON")
		return
	}

	startDate, err := time.Parse("2006-01-02", currentRequest.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error while parsing start Date")
		return
	}
	endDate, err := time.Parse("2006-01-02", currentRequest.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error while parsing end Date")
		return
	}

	excelFileName := "./resources/file/Date.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}

	total := 0

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) >= 2 {
				dateCell := row.Cells[0]
				numCell := row.Cells[1]

				dateStr := dateCell.String()
				date, err := time.Parse("2006-01-02", dateStr)
				if err != nil {
					continue
				}

				if date.After(startDate) && date.Before(endDate) || date.Equal(startDate) || date.Equal(endDate) {
					num, err := numCell.Int()
					if err != nil {
						fmt.Printf("Error parsing number: %s\n", err)
						continue
					}
					total += num
				}
			}
		}
	}

	c.JSON(http.StatusOK, total)
	return
}

type showImageRequest struct {
	StudentID string `json:"StudentID"`
}

func ShowImage(c *gin.Context) {
	var currentRequest showImageRequest
	if err := c.BindJSON(&currentRequest); err != nil {
		c.JSON(http.StatusBadRequest, "Error parsing JSON")
		return
	}

	identity := utils.MapStringToNumber(currentRequest.StudentID)
	response := gin.H{
		"identity": initialize.Configuration.StudentCategorise[identity],
		"image1":   initialize.Configuration.StudentImage[identity],
		"image2":   initialize.Configuration.StudentImage[6],
	}
	c.JSON(http.StatusOK, response)
	return
}

func UploadCSV(c *gin.Context) {
	// receive new file
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error while receiving the file: %s", err.Error()))
		return
	}
	defer file.Close()

	// 读取CSV内容
	csvFile, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 添加两列的标题
	csvFile[0] = append(csvFile[0], "Categorise")
	for i := 1; i < len(csvFile); i++ {
		identity := utils.MapStringToNumber(csvFile[i][0])
		csvFile[i] = append(csvFile[i], initialize.Configuration.StudentCategorise[identity])
	}

	// 将修改后的CSV数据写入响应中
	b := &bytes.Buffer{}
	w := csv.NewWriter(b)
	err = w.WriteAll(csvFile)
	if err != nil {
		log.Fatal(err)
	}

	c.Header("Content-Disposition", "attachment; filename=modified_"+header.Filename)
	c.Header("Content-Type", "text/csv")
	c.Data(http.StatusOK, "text/csv", b.Bytes())
}
