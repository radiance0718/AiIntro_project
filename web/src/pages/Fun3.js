import React, { useState } from 'react';
import { Input, Button, Row, Col } from 'antd';
import axios from 'axios';

function QueryStudent() {
  const [studentId, setStudentId] = useState('');
  const [queryResult, setQueryResult] = useState('');

  const handleInputChange = (e) => {
    setStudentId(e.target.value);
  };

  const handleQuery = async () => {
    try {
      // 替换为您的后端API地址和参数
      const response = await axios.get(`https://your-api.com/students/${studentId}`);
      setQueryResult(response.data); // 假设返回的数据在 response.data 中
    } catch (error) {
      console.error('查询失败:', error);
      setQueryResult('查询失败，请重试');
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
          {queryResult && <div>查询结果: {JSON.stringify(queryResult)}</div>}
        </Col>
      </Row>
    </div>
  );
}

export default QueryStudent;
