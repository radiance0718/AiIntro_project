import React, { useState } from 'react';
import { Select, Row, Col } from 'antd';

const { Option } = Select;

function Home() {
  const [selectedCategory, setSelectedCategory] = useState('');

  // 分类及对应人数
  const categoryCounts = {
    '勤学者': 120,
    '思考者': 80,
    '博览者': 150
  };

  const handleCategoryChange = value => {
    setSelectedCategory(value);
  };

  return (
    <div>
      <Row gutter={16} style={{ padding: 20 }}>
        <Col span={12}>
          <div className="section">
            <h2>选择分类</h2>
            <Select
              style={{ width: 200 }}
              placeholder="请选择一个分类"
              onChange={handleCategoryChange}
              value={selectedCategory}
            >
              {Object.keys(categoryCounts).map(category => (
                <Option key={category} value={category}>{category}</Option>
              ))}
            </Select>
            {selectedCategory && (
              <p>
                {selectedCategory}类有{categoryCounts[selectedCategory]}人
              </p>
            )}
          </div>
        </Col>
        <Col span={12}><div className="section">区域2</div></Col>
        <Col span={12}><div className="section">区域3</div></Col>
        <Col span={12}><div className="section">区域4</div></Col>
      </Row>
    </div>
  );
}

export default Home;
