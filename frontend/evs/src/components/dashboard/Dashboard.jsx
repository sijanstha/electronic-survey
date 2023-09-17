import React from "react";
import Sidebar from "../fragment/Sidebar";
import Navbar from "../fragment/Navbar";

const Dashboard = () => {
    return (
        <div id="viewport">
                <Sidebar />
                <Navbar />
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
    );
}

export default Dashboard;