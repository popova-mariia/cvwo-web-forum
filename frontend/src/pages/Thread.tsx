import React, { useEffect, useState } from 'react';
import axios from 'axios';

// Define a type for the thread
interface ThreadType {
  id: number;
  title: string;
}

const ThreadPage: React.FC = () => {
  // Use the defined type for the state
  const [threads, setThreads] = useState<ThreadType[]>([]);

  useEffect(() => {
    const fetchThreads = async () => {
      try {
        console.log("Fetching threads");
        const response = await axios.get('http://localhost:8080/threads', { withCredentials: true });        
        console.log("Response:", response);        
        setThreads(response.data);
        
      } catch (error) {
        console.error('Error fetching threads:', error);
      }
    };

    fetchThreads();
  }, []);

  return (
    <div>
      <h2>Forum Threads</h2>
      <ul>
        {threads.map((thread) => (
          <li key={thread.id}>{thread.title}</li>
        ))}
      </ul>
    </div>
  );
};

export default ThreadPage;