import React, { useState } from 'react';
import axios from 'axios';

const User: React.FC = () => {
  const [username, setUsername] = useState('');
  const [score, setScore] = useState(0);
  const [isAnonymous, setIsAnonymous] = useState(false);

  const createUser = async () => {
    try {
      const response = await axios.post('/users', {
        username,
        score,
        is_anonymous: isAnonymous,
      });
      console.log('User created:', response.data);
    } catch (error) {
      console.error('Error creating user:', error);
    }
  };

  return (
    <div>
      <h2>Create User</h2>
      <input
        type="text"
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <input
        type="number"
        placeholder="Score"
        value={score}
        onChange={(e) => setScore(Number(e.target.value))}
      />
      <label>
        <input
          type="checkbox"
          checked={isAnonymous}
          onChange={(e) => setIsAnonymous(e.target.checked)}
        />
        Anonymous
      </label>
      <button onClick={createUser}>Create User</button>
    </div>
  );
};

export default User;