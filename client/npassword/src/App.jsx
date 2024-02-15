// App.jsx
import React, { useState } from 'react';
import axios from 'axios';
import './App.css'; // Import CSS file for styling

function App() {
  const [userID, setUserID] = useState('');
  const [password, setPassword] = useState('');
  const [status, setStatus] = useState('');

  const handleCreatePassword = async () => {
    try {
      const response = await axios.post('http://localhost:8080/passwords', { id: userID, password });
      if (response.status === 201) {
        setStatus('Password created successfully.');
      }
    } catch (error) {
      setStatus(`Error: ${error.message}`);
    }
  };

  return (
    <div className="container">
      <h1>nPassword:</h1>
      <h2>Distributed Password Manager</h2>
      <div className="form-group">
        <label>User ID: </label>
        <input type="text" className="form-control" value={userID} onChange={(e) => setUserID(e.target.value)} />
      </div>
      <div className="form-group">
        <label>Password: </label>
        <input type="password" className="form-control" value={password} onChange={(e) => setPassword(e.target.value)} />
      </div>
      <button className="btn btn-primary" onClick={handleCreatePassword}>Add Password</button>
      <div className="status">{status}</div>
    </div>
  );
}

export default App;