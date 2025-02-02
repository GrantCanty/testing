import React from 'react'
import SidebarOption from "./sidebar-option";
import { useNavigate, Outlet } from "react-router-dom";
import './sidebar.css'

const Sidebar = ({homeRoute, deckRoute}) => {
    const navigate = useNavigate()

    return (
        <>
        <div className="sidebar-wrapper">
            <SidebarOption logo="🏠" title="Home" onClick={() => navigate(homeRoute)} />
            <SidebarOption logo="🗂"  title="Decks" onClick={() => navigate(deckRoute)} />
        </div>
        <div id="detail">
            <Outlet />
        </div>
        </>
    )
}

export default Sidebar;