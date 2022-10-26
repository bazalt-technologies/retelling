import CompRegistration from "@/components/LoginAndRegister/CompRegistration";
import LoginComponent from "@/components/LoginAndRegister/LoginComponent";
import MainView from "@/components/MainView/MainView";
import OpenComponent from "@/components/OpenComponent";
import ProfileSettings from "@/components/Settings/ProfileSettings";
import ProfileReviewsComponent from "@/components/Profile/ProfileReviewsComponent";
import ContentWithReviews from "@/components/MainView/ContentWithReviews";
import NewReviewComponent from "@/components/Profile/NewReviewComponent";

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
    {
        path: "/content/:id/reviews",
        name: "contentReviews",
        component: ContentWithReviews
    },
    {
        path: "/profile/newReview",
        name: "newReview",
        component: NewReviewComponent
    }

]

export default routes