import CompRegistration from "@/components/CompRegistration";
import LoginComponent from "@/components/LoginComponent";
import MainView from "@/components/MainView";
import OpenComponent from "@/components/OpenComponent";

const routes = [
    {
        path: "/",
        component: OpenComponent
    },
    {
        path: "/registration",
        component: CompRegistration
    },
    {
        path: "/login",
        component: LoginComponent
    },
    {
        path: "/content",
        component: MainView
    }
]

export default routes