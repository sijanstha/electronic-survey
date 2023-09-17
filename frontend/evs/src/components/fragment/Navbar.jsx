import React from "react";

const Navbar = () => {
    return (
        <React.Fragment>
            <nav className="navbar navbar-expand-lg bg-body-tertiary mb-5 p-2" style={{ "boxShadow": "rgba(0, 0, 0, 0.1) 0px 4px 12px" }}>
                <div className="container-fluid">
                    <img className="navbar-brand" src="/logo.png" width="50px" alt="" />
                    <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div className="collapse navbar-collapse" id="navbarSupportedContent">
                        <ul className="navbar-nav me-auto mb-2 mb-lg-0">

                        </ul>
                        <div className="d-flex list-style-none" style={{ "listStyle": "none" }}>
                            <li className="nav-item dropdown">
                                <a className="nav-link dropdown-toggle" href="" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                                    Email
                                </a>
                                <ul className="dropdown-menu">
                                    <li><a className="dropdown-item" href="/logout">Logout</a></li>
                                </ul>
                            </li>
                        </div>
                    </div>
                </div>
            </nav>
        </React.Fragment>
    );
};

export default Navbar;