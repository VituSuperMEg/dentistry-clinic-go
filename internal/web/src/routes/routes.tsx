import SignIn from "@/pages/auth/sign-in";
import SignUp from "@/pages/auth/sign-up";
import { Home } from "@/pages/dashboard/home";
import Notifications from "@/pages/dashboard/notifications";
import Profile from "@/pages/dashboard/profile";
import Tables from "@/pages/dashboard/tables";
import {
  HomeIcon,
  InfoIcon,
  RectangleVertical,
  StepBackIcon,
  TableCellsMergeIcon,
  UserCircleIcon,
} from "lucide-react";

const icon = {
  className: "w-5 h-5 text-inherit",
};

export const routes = [
  {
    layout: "dashboard",
    pages: [
      {
        icon: <HomeIcon {...icon} />,
        name: "dashboard",
        path: "/home",
        element: <Home />,
      },
      {
        icon: <UserCircleIcon {...icon} />,
        name: "Pacientes",
        path: "/profile",
        element: <Profile />,
      },
      {
        icon: <TableCellsMergeIcon {...icon} />,
        name: "tables",
        path: "/tables",
        element: <Tables />,
      },
      {
        icon: <InfoIcon {...icon} />,
        name: "notifications",
        path: "/notifications",
        element: <Notifications />,
      },
    ],
  },
  {
    title: "auth pages",
    layout: "auth",
    pages: [
      {
        icon: <StepBackIcon {...icon} />,
        name: "sign in",
        path: "/sign-in",
        element: <SignIn />,
      },
      {
        icon: <RectangleVertical {...icon} />,
        name: "sign up",
        path: "/sign-up",
        element: <SignUp />,
      },
    ],
  },
];

export default routes;
