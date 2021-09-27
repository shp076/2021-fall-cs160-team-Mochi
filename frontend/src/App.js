import './App.css';
import { useState } from 'react';

function App() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async(e) => {
    e.preventDefault();

    await fetch('http://localhost:3000/v1/login?signup=true', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify({
        username,
        password
      })
  });

  };

  return (
    <div className="App">
      <main className="form-signup">
        <form onSubmit={handleSubmit}>
            <h1 className="h3 mb-3 fw-normal">Please sign up</h1>
            <input type ="username" className="form-control" placeholder="Username" required
                   onChange={e => setUsername(e.target.value)}
            />
            <input type="password" className="form-control" placeholder="Password" required
                   onChange={e => setPassword(e.target.value)}
            />
            <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
        </form>
      </main>
    </div>
  );
}

export default App;
