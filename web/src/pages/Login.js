import React from 'react';
import { Form, Input, Button } from 'antd';
import { useNavigate } from 'react-router-dom';

function Login() {
  const [form] = Form.useForm();
  const navigate = useNavigate();

  const onFinish = async (values) => {
    try {
      const response = await fetch('http://10.26.137.106:9090/admin/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Name: values.username,
          Password: values.password,
        }),
      });
      const data = await response.json();
      console.log(data)

      if (data=="true") {
        // 登录成功
        console.log('Login successful');
        navigate('/home'); // 跳转到 home 页面
      } else {
        // 登录失败
        console.log('Login failed');
        alert('Invalid username or password');
      }
    } catch (error) {
      // 网络或其他错误
      console.error('Login error:', error);
      alert('Login error');
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
