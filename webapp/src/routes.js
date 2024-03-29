import CompRegistration from "@/components/LoginAndRegister/CompRegistration";
import LoginComponent from "@/components/LoginAndRegister/LoginComponent";
import MainView from "@/components/MainView/MainView";
import OpenComponent from "@/components/OpenComponent";
import ProfileSettings from "@/components/Settings/SettingsScreen";
import ProfileReviewsComponent from "@/components/Profile/ProfileReviewsComponent";
import ContentWithReviews from "@/components/MainView/ContentWithReviews";
import SearchComponent from "@/components/MainView/SearchComponent";
import SettingsComponent from "@/components/Settings/SettingsComponent";
import RestContent from "@/components/MainView/RestContent";

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
        path: "/content/recommendations",
        component: MainView
    },
    {
        path: "/content/new",
        component: RestContent
    },
    {
        path: "/content/search",
        component: SearchComponent
    },
    {
        path: "/settings",
        component: ProfileSettings
    },
    {
        path: "/profile",
        component: ProfileReviewsComponent
    },
    {
        path: "/content/:id/reviews",
        name: "contentReviews",
        component: ContentWithReviews
    },
    {
        path: "/settings/profile",
        component: SettingsComponent
    }

]

export default routes