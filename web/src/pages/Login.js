import React from 'react';
import { Form, Input, Button } from 'antd';
import { useNavigate } from 'react-router-dom'; // 引入 React Router 的 useNavigate 钩子

function Login() {
  const [form] = Form.useForm();
  const navigate = useNavigate(); // 创建 navigate 函数实例

  // 预设的有效用户名和密码
  const validUsername = 'admin';
  const validPassword = '123';

  const onFinish = (values) => {
    // 验证输入的用户名和密码
    if (values.username === validUsername && values.password === validPassword) {
      // 登录成功
      console.log('Login successful');
      navigate('/home'); // 跳转到 dashboard 页面
    } else {
      // 登录失败
      console.log('Login failed');
      alert('Invalid username or password'); // 显示错误信息
    }
  };

  return (
    <div style={{ maxWidth: 300, margin: 'auto' }}>
      <h2>Login</h2>
      <Form
        form={form}
        name="login_form"
        onFinish={onFinish}
      >
        <Form.Item
          name="username"
          rules={[{ required: true, message: 'Please input your Username!' }]}
        >
          <Input placeholder="Username" />
        </Form.Item>
        <Form.Item
          name="password"
          rules={[{ required: true, message: 'Please input your Password!' }]}
        >
          <Input.Password placeholder="Password" />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" style={{ width: '100%' }}>
            Log in
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
}

export default Login;
