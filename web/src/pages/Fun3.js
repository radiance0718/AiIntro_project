import React, { useState } from 'react';
import { Input, Button, Row, Col, Upload, message } from 'antd';
import axios from 'axios';

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
      setCategoryImages(response.data.images);
    } catch (error) {
      console.error('图片获取失败:', error);
      message.error('图片获取失败，请重试');
    }
  };

  // 处理文件变化
  const handleFileChange = ({ file }) => {
    // 直接设置文件数据
    setFileData(file);
  };

  // 上传文件
  const handleFileUpload = async () => {
    try {
      const formData = new FormData();
      formData.append('file', fileData); // 上传原始文件对象

      const response = await axios.post('http://10.26.137.106:9090/admin/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
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
          {categoryImages.map((img, index) => (
            <img key={index} src={`data:image/png;base64,${img}`} alt={`category-${index}`} />
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
