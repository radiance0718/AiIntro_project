import React, { useState } from 'react';
import { Input, Button, Row, Col, Upload, message, Image } from 'antd';
import axios from 'axios';
import Papa from 'papaparse';

function QueryStudent() {
  const [studentId, setStudentId] = useState('');
  const [queryResult, setQueryResult] = useState('');
  const [categoryImages, setCategoryImages] = useState([]);
  const [fileData, setFileData] = useState(null);

  // 处理输入框变化
  const handleInputChange = (e) => {
    setStudentId(e.target.value);
  };

  // 查询学号
  const handleQuery = async () => {
    try {
      const response = await axios.get(`https://your-api.com/students/${studentId}`);
      setQueryResult(response.data);
      // 假设类别是响应数据的一部分
      const category = response.data.category;
      fetchCategoryImages(category);
    } catch (error) {
      console.error('查询失败:', error);
      setQueryResult('查询失败，请重试');
    }
  };

  // 获取类别图片
  const fetchCategoryImages = async (category) => {
    try {
      const response = await axios.get(`https://your-api.com/category-images/${category}`);
      // 假设响应包含一个base64图片数组
      setCategoryImages(response.data.images);
    } catch (error) {
      console.error('图片获取失败:', error);
      message.error('图片获取失败，请重试');
    }
  };

  // 处理文件变化
  const handleFileChange = ({ file }) => {
    const reader = new FileReader();
    reader.onload = (e) => {
      Papa.parse(e.target.result, {
        complete: (results) => {
          setFileData(results.data);
          message.success('文件读取成功');
        },
        header: true
      });
    };
    reader.readAsText(file);
  };

  // 上传文件
  const handleFileUpload = async () => {
    try {
      const response = await axios.post('https://your-api.com/upload-csv', {
        data: fileData
      });
      message.success('文件上传成功');
    } catch (error) {
      console.error('文件上传失败:', error);
      message.error('文件上传失败，请重试');
    }
  };

  return (
    <div style={{ padding: '20px' }}>
      <Row gutter={16}>

        {/* 学生信息查询部分 */}
        <Col span={12}>
          <Input 
            placeholder="输入学号" 
            value={studentId} 
            onChange={handleInputChange} 
          />
          <Button type="primary" onClick={handleQuery}>查询</Button>
          {queryResult && <div>查询结果: {JSON.stringify(queryResult)}</div>}
          {/* 显示类别图片 */}
          {categoryImages.map((img, index) => (
            <Image key={index} src={`data:image/png;base64,${img}`} />
          ))}
        </Col>

        {/* CSV文件处理部分 */}
        <Col span={12}>
          <Upload 
            beforeUpload={() => false} 
            onChange={handleFileChange}>
            <Button>选择CSV文件</Button>
          </Upload>
          <Button 
            type="primary" 
            onClick={handleFileUpload} 
            disabled={!fileData}
          >
            上传文件
          </Button>
        </Col>

      </Row>
    </div>
  );
}

export default QueryStudent;
