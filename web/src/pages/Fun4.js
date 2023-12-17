import React, { useState } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import './Fun4.css'; // 导入样式文件

const DateRangePicker = ({ startDate, endDate, handleStartDateChange, handleEndDateChange }) => {
    return (
        <div className="date-range-picker">
            <label>开始日期：</label>
            <DatePicker selected={startDate} onChange={handleStartDateChange} />
            <br />
            <label>结束日期：</label>
            <DatePicker selected={endDate} onChange={handleEndDateChange} />
        </div>
    );
};

const VisitStatistics = ({ startDate, endDate }) => {
    const fetchVisitCount = (startDate, endDate) => {
        // 这里可以是同步的数据获取逻辑
        // 返回一个假设的访问数量
        return Math.floor(Math.random() * 100);
    };

    const displayMessage = () => {
        try {
            const visitCount = fetchVisitCount(startDate, endDate);
            return `在此期间有 ${visitCount} 个人访问了图书馆分类系统。`;
        } catch (error) {
            console.error('Error fetching visit count:', error);
            return '无法获取访问统计信息。';
        }
    };

    return (
        <div className="visit-statistics">
            <p>{displayMessage()}</p>
        </div>
    );
};

function Fun4() {
    const [startDate, setStartDate] = useState(new Date());
    const [endDate, setEndDate] = useState(new Date());

    const handleStartDateChange = (date) => {
        setStartDate(date);
        if (endDate < date) {
            setEndDate(date);
        }
    };

    const handleEndDateChange = (date) => {
        if (date >= startDate) {
            setEndDate(date);
        }
    };

    return (
        <div className="fun4-container">
            <h2>图书馆分类系统访问统计</h2>
            <DateRangePicker
                startDate={startDate}
                endDate={endDate}
                handleStartDateChange={handleStartDateChange}
                handleEndDateChange={handleEndDateChange}
            />
            <VisitStatistics startDate={startDate} endDate={endDate} />
        </div>
    );
}

export default Fun4;
