import React from 'react';

const ThreadPage: React.FC = () => {
    const username = localStorage.getItem('username');

    return (
        <div>
            <h1>This is a thread page</h1>
            <p>Your username is: {username}</p>
        </div>
    );
};

export default ThreadPage;