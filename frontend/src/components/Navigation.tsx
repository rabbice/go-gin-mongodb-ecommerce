import React from 'react'
import { Link } from 'react-router-dom';

const Nav = () => {
    return (
        <div>
               <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
  <div className="container-fluid">
    <Link to={"/"} className="navbar-brand">Marketplace Service</Link>
    <div>
      <ul className="navbar-nav me-auto mb-2 mb-md-0">
      <li className="nav-item-shop">
          <Link to={"/shops"} className="nav-link active" aria-current="page">Shops</Link>
        </li>
        <li className="nav-item">
          <Link to={"/login"} className="nav-link active" aria-current="page">Login</Link>
        </li>
        <li className="nav-item">
          <Link to={"/signup"} className="nav-link active" aria-current="page">Register</Link>
        </li>
      </ul>
    </div>
  </div>
</nav>
      
            </div>
    );
};

export default Nav;