import CompRegistration from "@/components/CompRegistration";
import LoginComponent from "@/components/LoginComponent";
import MainView from "@/components/MainView";
import OpenComponent from "@/components/OpenComponent";
import ProfileSettings from "@/components/ProfileSettings";
import ProfileReviewsComponent from "@/components/ProfileReviewsComponent";

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
    },
    {
        path: "/settings",
        component: ProfileSettings
    },
    {
        path: "/profile",
        component: ProfileReviewsComponent
    },

]

export default routes