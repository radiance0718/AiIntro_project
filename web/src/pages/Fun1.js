import React, { useState } from 'react';
import { Select, Button, Row, Col } from 'antd';
import axios from 'axios';

const { Option } = Select;

function Fun1() {
  const [selectedCategory, setSelectedCategory] = useState('');
  const [queryResult, setQueryResult] = useState('');

  // 使用对象数组来定义分类和它们的汉字表示
  const categories = [
    { label: '阅览者', value: '1' },
    { label: '钻研者', value: '2' },
    { label: '思辨者', value: '3' },
    { label: '勤学者', value: '4' },
    { label: '学习者', value: '5' },
    { label: '无闻者', value: '6' }
    // ...可以根据需要添加更多分类
  ];

  const handleCategoryChange = value => {
    setSelectedCategory(value);
  };

  const handleQuery = async () => {
    try {
      const response = await axios.post('http://10.26.137.106:9090/admin/countType', {
        Type: selectedCategory
      });
      setQueryResult(response.data);
      console.log(response.data);
    } catch (error) {
      console.error('查询失败:', error);
    }
  };

  const centerStyle = {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
   
  };

  return (
    <div style={centerStyle}>
      <Row gutter={16}>
        <Col>
          <Select value={selectedCategory} onChange={handleCategoryChange} style={{ width: 200 }}>
            {categories.map(category => (
              <Option key={category.value} value={category.value}>{category.label}</Option>
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
