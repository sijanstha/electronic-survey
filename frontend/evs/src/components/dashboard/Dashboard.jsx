import React from "react";

const Dashboard = () => {

    return (
        <React.Fragment>
            <div id="viewport">
                <div id="sidebar">
                    <header>
                        <a href="#">Electronic Voting System</a>
                    </header>
                    <ul class="nav">
                        <li>
                            <a href="/">
                                <i class="zmdi zmdi-view-dashboard"></i> Dashboard
                            </a>
                        </li>
                        <li>
                            <a href="/admin/organizers">
                                <i class="zmdi zmdi-link"></i> Manage Organizers
                            </a>
                        </li>
                        <li>
                            <a href="/poll">
                                <i class="zmdi zmdi-calendar"></i> Manage Polls
                            </a>
                        </li>
                    </ul>
                </div>
                <nav class="navbar navbar-expand-lg bg-body-tertiary mb-5 p-2" style={{ "box-shadow": "rgba(0, 0, 0, 0.1) 0px 4px 12px;" }}>
                    <div class="container-fluid">
                        <img class="navbar-brand" src="/logo.png" width="50px" alt="" />
                        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                            <span class="navbar-toggler-icon"></span>
                        </button>
                        <div class="collapse navbar-collapse" id="navbarSupportedContent">
                            <ul class="navbar-nav me-auto mb-2 mb-lg-0">

                            </ul>
                            <div class="d-flex list-style-none" style={{ "list-style": "none" }}>
                                <li class="nav-item dropdown">
                                    <a class="nav-link dropdown-toggle" href="" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                                        Email
                                    </a>
                                    <ul class="dropdown-menu">
                                        <li><a class="dropdown-item" href="/logout">Logout</a></li>
                                    </ul>
                                </li>
                            </div>
                        </div>
                    </div>
                </nav>
                <div class="container">
                    <div id="content">
                        <div class="container-fluid">
                            <h4>Simple Sidebar</h4>
                            <p>
                                Make sure to keep all page content within the
                                <code>#content</code>.
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </React.Fragment>
    );
}

export default Dashboard;