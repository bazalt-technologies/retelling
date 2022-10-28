import CompRegistration from "@/components/LoginAndRegister/CompRegistration";
import LoginComponent from "@/components/LoginAndRegister/LoginComponent";
import MainView from "@/components/MainView/MainView";
import OpenComponent from "@/components/OpenComponent";
import ProfileSettings from "@/components/Settings/ProfileSettings";
import ProfileReviewsComponent from "@/components/Profile/ProfileReviewsComponent";
import ContentWithReviews from "@/components/MainView/ContentWithReviews";
import NewReviewComponent from "@/components/Profile/NewReviewComponent";
import SearchComponent from "@/components/MainView/SearchComponent";
import SettingsComponent from "@/components/Settings/SettingsComponent";

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
        path: "/profile/newReview",
        name: "newReview",
        component: NewReviewComponent
    },
    {
        path: "/settings/profile",
        component: SettingsComponent
    }

]

export default routes