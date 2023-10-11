import React from "react";
import "./sidebar.css";
const Sidebar = () => {
  return (
    <React.Fragment>
        <div id="sidebar">
                    <header>
                        <a href="/">Electronic Voting System</a>
                    </header>
                    <ul className="nav">
                        <li>
                            <a href="/">
                                <i className="zmdi zmdi-view-dashboard"></i> Dashboard
                            </a>
                        </li>
                        <li>
                            <a href="/admin/organizers">
                                <i className="zmdi zmdi-link"></i> Manage Organizers
                            </a>
                        </li>
                        <li>
                            <a href="/poll/add">
                                <i className="zmdi zmdi-calendar"></i> Add Poll
                            </a>
                        </li>
                        <li>
                            <a href="/poll">
                                <i className="zmdi zmdi-calendar"></i> Manage Polls
                            </a>
                        </li>
                        <li>
                            <a href="/participant-list">
                                <i className="zmdi zmdi-calendar"></i> Manage Participant List
                            </a>
                        </li>
                    </ul>
                </div>
    </React.Fragment>
  );
};

export default Sidebar;