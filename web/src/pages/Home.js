import React from 'react';
import { Button, Row, Col } from 'antd';
import { useNavigate } from 'react-router-dom';

function Home() {
  const navigate = useNavigate();

  const handleNavigate = (path) => {
    navigate(path);
  };

  const buttonStyle = (bgColor) => ({
    width: '100%',
    height: '50vh',
    fontSize: '18px',
    backgroundColor: bgColor, // 背景颜色
    border: 'none',           // 移除边框
    color: 'black'            // 字体颜色
  });

  return (
    <Row style={{ height: '100vh' }}>
      <Col span={12} style={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
        <Button style={buttonStyle('#fadb5f')} onClick={() => handleNavigate('/Fun1')}>查询各个分类人数</Button> {/* 淡黄色 */}
      </Col>
      <Col span={12} style={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
        <Button style={buttonStyle('#95de64')} onClick={() => handleNavigate('/Fun2')}>查询各个年级各分类人数</Button> {/* 淡绿色 */}
      </Col>
      <Col span={12} style={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
        <Button style={buttonStyle('#69c0ff')} onClick={() => handleNavigate('/Fun3')}>获取指定用户分类以及画像</Button> {/* 淡蓝色 */}
      </Col>
      <Col span={12} style={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
        <Button style={buttonStyle('#ff85c0')} onClick={() => handleNavigate('/Fun4')}>图书馆分类系统访问统计</Button> {/* 淡粉色 */}
      </Col>
    </Row>
  );
}

export default Home;
