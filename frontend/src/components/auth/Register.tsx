import React, {SyntheticEvent, useState} from 'react'
import {Navigate} from 'react-router-dom'

const SignUp = () => {
    const [firstname, setFirstName] = useState('');
    const [lastname, setLastName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [navigate, setNavigate] = useState(false);
    const [seller, setSeller] = useState(false);
    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        await fetch("http://localhost:5000/auth/signup", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify({
                    firstname,
                    lastname,
                    email,
                    password,
                    seller
            })

        });
        setNavigate(true);
    }

    if (navigate) {
        return <Navigate to="/login" />
    }

  
    
    return (
        <main className="form-signin">
            <form onSubmit={submit}>
                <h1 className="h3 mb-3 fw-normal">Please Register</h1>
                    <input type="firstname" className="form-control" placeholder="First name" 
                        onChange={e => setFirstName(e.target.value)}
                    />
                    <input type="lastname" className="form-control" placeholder="Last name" 
                        onChange={e => setLastName(e.target.value)}
                    />
                    <input type="email" className="form-control" placeholder="Email" 
                        onChange={e => setEmail(e.target.value)}
                    />
                    <input type="password" className="form-control" id="floatingPassword" placeholder="Password" 
                        onChange={e => setPassword(e.target.value)}
                    />
                    <div className="form-check form-check-inline">
                    <input className="form-check-input" type="radio" name="seller" value="option1" 
                    />
                        Buyer
                    </div>
                    <div className="form-check form-check-inline">
                    <input className="form-check-input" type="radio" name="seller" value="option2" 
                        onChange={e => setSeller(e.target.checked)}
                    />
                        Seller
                    </div>
                    <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
                    <p className="mt-5 mb-3 text-muted">&copy; 2017–2021</p>
            </form>
        </main>
    );
};

export default SignUp;