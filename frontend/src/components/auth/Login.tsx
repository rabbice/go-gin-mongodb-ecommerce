import React, {SyntheticEvent, useState} from 'react'
import { Navigate } from 'react-router-dom'


const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const[navigate, setNavigate] = useState(false);
    

    const submit = async (e: SyntheticEvent) => {
      e.preventDefault();

      await fetch("http://localhost:5000/auth/login", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        credentials: "include",
        body: JSON.stringify({
                email,
                password
        })
    });
    setNavigate(true);
    }
    if (navigate) {
      return <Navigate to="/login" />
  }
    return (
        <div>
           <main className="form-signin">
            <form onSubmit={submit}>
    <h1 className="h3 mb-3 fw-normal">Please sign in</h1>

      <input type="email" className="form-control" id="floatingInput" placeholder="name@example.com" 
        onChange={e => setEmail(e.target.value)}
      />
      <input type="password" className="form-control" id="floatingPassword" placeholder="Password" 
        onChange={e => setPassword(e.target.value)}
      />

    <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
    <p className="mt-5 mb-3 text-muted">&copy; 2017–2021</p>
  </form>
</main>
        </div>
    );
};

export default Login;