import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { useAuth } from "../provider/authProvider";
import { ProtectedRoute } from "./ProtectedRoute";
import Login from "../components/login/Login";
import Poll from "../components/poll/Poll";
import Register from "../components/register/Register";
import Dashboard from "../components/dashboard/Dashboard";
import AddPoll from "../components/poll/AddPoll";
import UpdatePoll from "../components/poll/UpdatePoll";

// https://dev.to/sanjayttg/jwt-authentication-in-react-with-react-router-1d03
const Routes = () => {
  const { token } = useAuth();

  // Define public routes accessible to all users
  const routesForPublic = [
    {
      path: "/register",
      element: <Register />,
    },
    {
      path: "/about-us",
      element: <div>About Us</div>,
    },
    {
      path: "/login",
      element: <Login />,
    },
  ];

  // Define routes accessible only to authenticated users
  const routesForAuthenticatedOnly = [
    {
      path: "/",
      element: <ProtectedRoute />, // Wrap the component in ProtectedRoute
      children: [
        {
          path: "/",
          element: <Dashboard />,
        },
        {
          path: "/poll",
          element: <Poll />,
        },
        {
          path: "/poll/add",
          element: <AddPoll />,
        },
        {
          path: "/poll/edit",
          element: <UpdatePoll />,
        },
        {
          path: "/poll/edit/:id",
          element: <UpdatePoll />,
        },
        {
          path: "/admin",
          element: <div>Admin Page</div>,
        },
      ],
    },
  ];

  // Define routes accessible only to non-authenticated users
  const routesForNotAuthenticatedOnly = [
    {
      path: "/login",
      element: <Login />,
    },
  ];

  // Combine and conditionally include routes based on authentication status
  const router = createBrowserRouter([
    ...routesForPublic,
    ...(!token ? routesForNotAuthenticatedOnly : []),
    ...routesForAuthenticatedOnly,
  ]);

  // Provide the router configuration using RouterProvider
  return <RouterProvider router={router} />;
};

export default Routes;