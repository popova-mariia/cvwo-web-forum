import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import '../style/Login.css';

const Login = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            console.log("Login submitted");
            const response = await axios.post('http://localhost:8080/login', { username, password });
            console.log('Login successful:', response.data);
            // Store the username in local storage or context
            localStorage.setItem('username', username);
            // Redirect to the thread page
            navigate('/threads');
        } catch (error) {
            if (axios.isAxiosError(error)) {
                console.error('Error logging in:', error.response?.data || error.message);
                console.error('Error details:', error.toJSON());
            } else {
                console.error('Unexpected error:', error);
            }
        }
    };

    return (
        <main className='login'>
            <form className='loginForm' onSubmit={handleSubmit}>
                <label htmlFor='username'>Username</label>
                <input
                    type='text'
                    name='username'
                    id='username'
                    required
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />
                <label htmlFor='password'>Password</label>
                <input
                    type='password'
                    name='password'
                    id='password'
                    required
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <button className='loginBtn'>SIGN IN</button>
                <p>
                    Don't have an account?{" "}
                    <a href='/register'>Create one</a>
                </p>
            </form>
        </main>
    );
};

export default Login;
