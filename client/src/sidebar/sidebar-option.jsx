import './sidebar-option.css'

const SidebarOption = ({logo, title, onClick}) => {
    return(
        <div className="sidebar-option" onClick={onClick}>
            <h1>{logo}</h1>
            <b><span className="sidebar-option-title">{title}</span></b>
        </div>
    )
}

export default SidebarOption;