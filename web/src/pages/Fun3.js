import React, { useState } from 'react';
import { Input, Button, Row, Col, Upload, message } from 'antd';
import axios from 'axios';


function QueryStudent() {
  const [studentId, setStudentId] = useState('');
  const [imageType, setImageType] = useState('');
  const [fileData, setFileData] = useState(null);

  // 处理学号输入框变化
  const handleInputChange = (e) => {
    setStudentId(e.target.value);
  };

  // 查询学号并获取图片类型
  const handleQuery = async () => {
    try {
      const response = await axios.post(`http://localhost:9090/admin/showImage`, {
        StudentID: studentId
      });
      setImageType(response.data.identity);
    } catch (error) {
      console.error('查询失败:', error);
      message.error('查询失败，请重试');
    }
  };

  // 处理文件选择变化
  const handleFileChange = ({ file }) => {
    setFileData(file);
  };

  const imageUrl = process.env.PUBLIC_URL + '/resource/Diligent.png';

  // 上传文件并下载返回的 CSV 文件
  const handleFileUpload = async () => {
    try {
      const formData = new FormData();
      formData.append('file', fileData); // 上传原始文件对象

      const response = await axios.post('http://localhost:9090/admin/upload', formData, {
        responseType: 'blob', // 告诉 axios 返回类型是文件
      });

      // 创建用于下载文件的链接
      const url = window.URL.createObjectURL(new Blob([response.data]));
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download', 'downloaded_file.csv'); // 设置下载文件名
      document.body.appendChild(link);
      link.click();

      message.success('文件上传成功，下载即将开始');
    } catch (error) {
      console.error('文件上传失败:', error);
      message.error('文件上传失败，请重试');
    }
  };

  // 获取图片路径
  const getImagePath = () => {
    return `/resource/${imageType}.png`; // 根据实际情况调整路径
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
      {imageType && (
            <img 
              src={getImagePath()} 
              alt="student-type" 
              width="500" 
              height="500"
            />)}
            <img 
              src="/resource/white.png"
              alt="student-type" 
              width="500" 
              height="500"
            />
    </div>
  );
}

export default QueryStudent;