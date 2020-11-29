import  MainPageContainer from "../Components/MainPage"
import  WidgetContainer from "../Components/Widget"
import EditorContainer from "../Components/Editor"

const stringToComp = {
    MainPageContainer: MainPageContainer,
    WidgetContainer: WidgetContainer,
    EditorContainer: EditorContainer
}

const ParseRoutes = (routes) => (
    routes.map((props) => ({...props, component: stringToComp[props.component]}))
)

export default ParseRoutes