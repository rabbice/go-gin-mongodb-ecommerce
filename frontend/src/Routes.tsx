import React from 'react'
import {Routes, Route} from 'react-router-dom'
import Login from './components/auth/Login';
import Home from './components/Home';
import NewShop from './components/shop/NewShop';
import SignUp from './components/auth/Register';
import AllShops from './components/shop/AllShops';

const MainRoutes = () => {
    return (
        <Routes>
              {/* public routes */}
        <Route path="/" element={<Home />} /> 
          <Route path="/login" element={<Login/>}/>
          <Route path="/signup" element={<SignUp/>}/>
          <Route path="/shops" element={<AllShops/>}/>
        {/* private routes */}
          <Route path="/create/shop" element={<NewShop />}/>
        </Routes>
    );
}

export default MainRoutes;