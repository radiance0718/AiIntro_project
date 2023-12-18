import React, { useState } from 'react';
import { Select, Button, Row, Col } from 'antd';
import axios from 'axios';

const { Option } = Select;

function Fun1() {
  const [selectedCategory, setSelectedCategory] = useState('');
  const [queryResult, setQueryResult] = useState('');

  const categories = ['1', '2', '3','4','5','6']; // 示例分类

  const handleCategoryChange = value => {
    setSelectedCategory(value);
  };

  const handleQuery = async () => {
    try {
        const response = await axios.post('http://10.26.137.106:9090/admin/countGrade', {
            Type: selectedCategory
          });
          


      setQueryResult(response.data);
       // 假设返回的数据在 response.data 中
       console.log(response.data)
    } catch (error) {
      console.error('查询失败:', error);
      // 处理错误情况
    }
  };

  return (
    <div style={{ padding: '20px' }}>
      <Row gutter={16}>
        <Col>
          <Select value={selectedCategory} onChange={handleCategoryChange} style={{ width: 200 }}>
            {categories.map(category => (
              <Option key={category} value={category}>{category}</Option>
            ))}
          </Select>
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

export default Fun1;
