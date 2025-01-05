import React, { useState, useEffect } from 'react';
import axios from 'axios';

// Define a type for the thread
interface ThreadType {
  id: number;
  title: string;
}

const Thread: React.FC = () => {
  // Use the defined type for the state
  const [threads, setThreads] = useState<ThreadType[]>([]);

  useEffect(() => {
    const fetchThreads = async () => {
      try {
        const response = await axios.get('/api/threads');
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

export default Thread;