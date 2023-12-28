import React, { useState, useEffect } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import axios from 'axios';
import './Fun4.css';

const formatDateToString = (date) => {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are zero-based
    const day = String(date.getDate()).padStart(2, '0');

    return `${year}-${month}-${day}`;
};

const DateRangePicker = ({ startDate, endDate, handleStartDateChange, handleEndDateChange, handleQuery }) => {
    return (
        <div className="date-range-picker">
            <label>开始日期：</label>
            <DatePicker selected={startDate} onChange={handleStartDateChange} />
            <br />
            <label>结束日期：</label>
            <DatePicker selected={endDate} onChange={handleEndDateChange} />
            <br />
            <button onClick={handleQuery}>查询</button>
        </div>
    );
};

const VisitStatistics = ({ startDate, endDate, queryButtonClicked, setQueryButtonClicked}) => {
    const [visitCount, setVisitCount] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            if (queryButtonClicked) {
                try {
                    const response = await axios.post(
                        'http://10.26.137.106:9090/admin/countDate',
                        {
                            StartDate: formatDateToString(startDate),
                            EndDate: formatDateToString(endDate)
                        }
                    );
                    setVisitCount(response.data);
                } catch (error) {
                    console.error('Error fetching visit count:', error);
                    setVisitCount(114);
                } finally {
                    // Reset queryButtonClicked to false after fetching data
                    setQueryButtonClicked(false);
                }
            } else{
                setVisitCount(514);
            }
        };

        fetchData();
    }, [startDate, endDate, queryButtonClicked, visitCount]);


    return (
        <div className="visit-statistics">
            <p>{visitCount == null ? '加载中...' : `在此期间有 ${visitCount} 个人访问了图书馆读者分类系统。`}</p>
        </div>
    );
};


function Fun4() {
    const [startDate, setStartDate] = useState(new Date());
    const [endDate, setEndDate] = useState(new Date());
    const [queryButtonClicked, setQueryButtonClicked] = useState(false);

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

    const handleQuery = () => {
        setQueryButtonClicked(true);
    };



    return (
        <div className="fun4-container">
            <h2>图书馆读者分类系统访问统计</h2>
            <DateRangePicker
                startDate={startDate}
                endDate={endDate}
                handleStartDateChange={handleStartDateChange}
                handleEndDateChange={handleEndDateChange}
                handleQuery={handleQuery}
            />
            {queryButtonClicked && (
                <VisitStatistics startDate={startDate}
                                 endDate={endDate}
                                 queryButtonClicked={queryButtonClicked}
                                 setQueryButtonClicked={setQueryButtonClicked}
                />

            )}
        </div>
    );
}

export default Fun4;
