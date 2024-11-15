import { Link, NavLink } from "react-router-dom";
import { IconButton } from "@material-tailwind/react";
import { BookMarked } from "lucide-react";
import {
  setOpenSidenav,
  useMaterialTailwindController,
} from "@/template/context";
import { Button } from "@/components/ui/button";

type SideNavType = {
  brandImg?: string; // Torne o tipo opcional
  brandName?: string; // Torne o tipo opcional
  routes: Array<{ path: string; name: string; icon: string }>;
};

export function Sidenav({
  brandImg = "/img/logo-ct.png",
  brandName = "Dentistry Clinic",
  routes,
}: SideNavType) {
  const [controller, dispatch] = useMaterialTailwindController();
  const { sidenavColor, sidenavType, openSidenav } = controller;

  const validColors = [
    "white",
    "black",
    "blue-gray",
    "gray",
    "brown",
    "deep-orange",
    "orange",
    "amber",
    "yellow",
    "lime",
    "light-green",
    "green",
    "teal",
    "cyan",
    "light-blue",
    "blue",
    "indigo",
    "deep-purple",
    "purple",
    "pink",
    "red",
  ];

  // Garantir que sidenavColor seja um valor válido
  const validSidenavColor = validColors.includes(sidenavColor)
    ? sidenavColor
    : "blue-gray";

  const sidenavTypes = {
    dark: "bg-gradient-to-br from-gray-800 to-gray-900",
    white: "bg-white shadow-sm",
    transparent: "bg-transparent",
  };

  return (
    <aside
      className={`${sidenavTypes[sidenavType]} ${
        openSidenav ? "translate-x-0" : "-translate-x-80"
      } fixed inset-0 z-50 my-4 ml-4 h-[calc(100vh-32px)] w-72 rounded-xl transition-transform duration-300 xl:translate-x-0 border border-blue-gray-100`}
    >
      <div className={`relative`}>
        <Link to="/" className="py-6 px-8 text-center">
          <h6
            color={sidenavType === "dark" ? "white" : "blue-gray"}
          >
            {brandName}
          </h6>
        </Link>
        <IconButton
          variant="text"
          color="white"
          size="sm"
          ripple={false}
          className="absolute right-0 top-0 grid rounded-br-none rounded-tl-none xl:hidden"
          onClick={() => setOpenSidenav(dispatch, false)}
        >
          <BookMarked strokeWidth={2.5} className="h-5 w-5 text-white" />
        </IconButton>
      </div>
      <div className="m-4">
        {routes.map(({ layout, title, pages }, key) => (
          <ul key={key} className="mb-4 flex flex-col gap-1">
            {title && (
              <li className="mx-3.5 mt-4 mb-2">
                <small
                  color={sidenavType === "dark" ? "white" : "blue-gray"}
                  className="font-black uppercase opacity-75"
                >
                  {title}
                </small>
              </li>
            )}
            {pages.map(({ icon, name, path }) => (
              <li key={name}>
                <NavLink to={`/${layout}${path}`}>
                  {({ isActive }) => (
                    <Button
                      variant={isActive ? "primary" : "ghost"}
                      color={
                        isActive
                          ? validSidenavColor 
                          : sidenavType === "dark"
                          ? "red"
                          : "blue-gray"
                      }
                      className="w-full flex justify-start"
                    >
                      {icon}
                      <p
                        color="inherit"
                        className="font-medium capitalize"
                      >
                        {name}
                      </p>
                    </Button>
                  )}
                </NavLink>
              </li>
            ))}
          </ul>
        ))}
      </div>
    </aside>
  );
}

export default Sidenav;
