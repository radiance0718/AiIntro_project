package services

import (
	"Demo/initialize"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"hash/fnv"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

var studentType map[string]int

type loginRequest struct {
	Name     string `json:"Name"`
	PassWord string `json:"PassWord"`
}

func Login(c *gin.Context) {
	var currentRequest loginRequest
	if err := c.ShouldBindJSON(&currentRequest); err != nil {
		c.JSON(http.StatusBadRequest, err)
		// fmt.Println("Error while parsing request: ", err.Error())
		return
	}

	adminName := currentRequest.Name
	if adminName != initialize.Configuration.AdminName {
		c.JSON(http.StatusBadRequest, "Wrong Admin Name")
		// fmt.Println("Wrong Admin Name")
		return
	}

	adminPassWord := currentRequest.PassWord
	if adminPassWord != initialize.Configuration.AdminPassword {
		c.JSON(http.StatusBadRequest, "Wrong Admin PassWord")
		// fmt.Println("Wrong admin Password")
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
		// fmt.Println("Error while parsing request: ", err.Error())
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
	// startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC) // 指定开始日期
	// endDate := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC) // 指定结束日期

	excelFileName := "./resources/file/Date.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}

	total := 0 // 存储总和的变量

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) >= 2 {
				dateCell := row.Cells[0]
				numCell := row.Cells[1]

				// 解析日期
				dateStr := dateCell.String()
				date, err := time.Parse("2006-01-02", dateStr)
				if err != nil {
					//fmt.Printf("Error parsing date: %s\n", err)
					continue
				}

				// 检查日期是否在指定范围内
				if date.After(startDate) && date.Before(endDate) || date.Equal(startDate) || date.Equal(endDate) {
					// 解析数字
					// _ := numCell.String()
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

	// fmt.Printf("Total sum of numbers within specified date range: %d\n", total)
	c.JSON(http.StatusOK, total)
	return
}

type checkTypeRequest struct {
	StudentID string `json:"StudentID"`
}

func CheckType(c *gin.Context) {
	var currentRequest checkTypeRequest
	if err := c.BindJSON(&currentRequest); err != nil {
		c.JSON(http.StatusBadRequest, "Error parsing JSON")
		return
	}

	if stuType, ok := studentType[currentRequest.StudentID]; ok {
		c.JSON(http.StatusOK, stuType)
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(6)
	studentType[currentRequest.StudentID] = randomNumber
	c.JSON(http.StatusOK, randomNumber)
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

	identity := mapStringToNumber(currentRequest.StudentID)
	image1Info := "./resources/image/Diligent.png"
	image1Data, err := ioutil.ReadFile(image1Info)
	if err != nil {
		log.Fatal(err)
	}
	image1Base64 := base64.StdEncoding.EncodeToString(image1Data)
	response := gin.H{
		"identity": identity,
		"image":    image1Base64,
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

	// create local file
	out, err := os.Create("./resources/file/" + header.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write file
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	// 打开CSV文件进行修改
	csvFile, err := os.OpenFile("./resources/file/"+header.Filename, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	// 读取CSV内容
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 添加两列的标题
	records[0] = append(records[0], "Categorise")
	for i := 1; i < len(records); i++ {
		identity := mapStringToNumber(records[i][0])
		switch identity {
		case 0:
			records[i] = append(records[i], "Diligent")
		case 1:
			records[i] = append(records[i], "Explorer")
		case 2:
			records[i] = append(records[i], "Learner")
		case 3:
			records[i] = append(records[i], "Researcher")
		case 4:
			records[i] = append(records[i], "Thinker")
		case 5:
			records[i] = append(records[i], "Unknown")
		}
	}

	// Write modified CSV data to a new file
	newCSVFile, err := os.Create("./resources/file/modified_" + header.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer newCSVFile.Close()

	w := csv.NewWriter(newCSVFile)
	err = w.WriteAll(records)
	if err != nil {
		log.Fatal(err)
	}

	// Send the modified CSV file to the frontend
	c.Header("Content-Disposition", "attachment; filename=modified_"+header.Filename)
	c.Header("Content-Type", "text/csv")
	c.File("./resources/file/modified_" + header.Filename)
}

func mapStringToNumber(input string) int {
	hash := fnv.New32a()
	hash.Write([]byte(input))
	hashValue := hash.Sum32()

	// 取模运算将哈希值映射到0到5之间的数字
	return int(hashValue % 6)
}
