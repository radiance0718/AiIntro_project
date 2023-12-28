import React, { useState } from 'react';
import { Input, Button, Row, Col, Upload, message } from 'antd';
import axios from 'axios';
import Papa from 'papaparse';

function QueryStudent() {
  const [studentId, setStudentId] = useState('');
  const [queryResult, setQueryResult] = useState('');
  const [fileData, setFileData] = useState(null);

  const handleInputChange = (e) => {
    setStudentId(e.target.value);
  };

  const handleQuery = async () => {
    try {
      const response = await axios.get(`https://your-api.com/students/${studentId}`);
      setQueryResult(response.data);
    } catch (error) {
      console.error('查询失败:', error);
      setQueryResult('查询失败，请重试');
    }
  };

  const handleFileChange = ({ file }) => {
    const reader = new FileReader();
    reader.onload = (e) => {
      Papa.parse(e.target.result, {
        complete: (results) => {
          setFileData(results.data);
          message.success('文件读取成功');
        },
        header: true // 如果CSV文件包含标题行，请设置为true
      });
    };
    reader.readAsText(file);
  };

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
        <Col>
          <Input 
            placeholder="输入学号" 
            value={studentId} 
            onChange={handleInputChange} 
          />
        </Col>
        <Col>
          <Button type="primary" onClick={handleQuery}>查询</Button>
        </Col>
      </Row>
      <Row style={{ marginTop: '20px' }}>
        <Col>
          <Upload 
            beforeUpload={() => false} 
            onChange={handleFileChange}>
            <Button>选择CSV文件</Button>
          </Upload>
        </Col>
        <Col>
          <Button 
            type="primary" 
            onClick={handleFileUpload} 
            disabled={!fileData}
          >
            上传CSV
          </Button>
        </Col>
      </Row>
      <Row style={{ marginTop: '20px' }}>
        <Col>
          {queryResult && <div>查询结果: {JSON.stringify(queryResult)}</div>}
        </Col>
      </Row>
    </div>
  );
}

export default QueryStudent;
