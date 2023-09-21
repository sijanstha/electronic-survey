import React from "react";
import "./sidebar.css";
import { Link } from "react-router-dom";

const Sidebar = () => {
  return (
    <React.Fragment>
        <div id="sidebar">
                    <header>
                        <a href="/">Electronic Voting System</a>
                    </header>
                    <ul className="nav">
                        <li>
                            <Link to="/">
                                <i className="zmdi zmdi-view-dashboard"></i> Dashboard
                            </Link>
                        </li>
                        <li>
                            <Link to="/admin/organizers">
                                <i className="zmdi zmdi-link"></i> Manage Organizers
                            </Link>
                        </li>
                        <li>
                            <Link to="/poll/add">
                                <i className="zmdi zmdi-calendar"></i> Add Poll
                            </Link>
                        </li>
                        <li>
                            <Link to="/poll">
                                <i className="zmdi zmdi-calendar"></i> Manage Polls
                            </Link>
                        </li>
                    </ul>
                </div>
    </React.Fragment>
  );
};

export default Sidebar;