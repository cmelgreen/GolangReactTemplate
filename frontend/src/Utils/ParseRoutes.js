import  MainPageContainer from "../Components/MainPage"
import  WidgetContainer from "../Components/Widget"

const stringToComp = {
    MainPageContainer: MainPageContainer,
    WidgetContainer: WidgetContainer
}

const ParseRoutes = (routes) => (
    routes.map((props) => ({...props, component: stringToComp[props.component]}))
)

export default ParseRoutes