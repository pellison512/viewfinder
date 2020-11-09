import React from "react"
class Drawing extends React.Component {

    constructor(props) {
        super(props);
        this.handleClick = this.handleClick.bind(this);
        this.state = {isClicked: false};
        this.state = {windowData: ''};
        this.state = {clientData: ''};
    }

    async handleClick() {
        let toggle = !this.state.isClicked

        await this.getDrawingInfo();

        this.setState(state =>({
            isClicked: toggle
        }));
    }

    async getDrawingInfo() {
        var title = document.title;
        console.log('hotdog: '+title);
        const encoded = encodeURI(title);

        var root = document.getElementById('root');
        var rect = root.getBoundingClientRect();
        console.log(rect);
        this.setState({
            clientData: rect
        });

        fetch('http://localhost:8090/windows/'+encoded).then(res => res.text()).then(res =>{
            this.setState({
                windowData: res
            });
        });
    }

    render() {
        


        return (
            <div>
                <h1>Hello From Create React App!</h1>
                <p>I am in a React Component</p>
                <button onClick={this.handleClick}>
                    Click Here!
                </button>
                <div >
                    {this.state.isClicked ? "ON":"OFF"}
                </div>
                <p>
                    Client Rectangle: 
                    <br/>left: {this.state.clientData.x}
                    <br/>top: {this.state.clientData.top}
                    <br/>right: {this.state.clientData.right}
                    <br/>bottom: {this.state.clientData.bottom}
                </p>
                <p>
                    WindowData: {this.state.windowData}
                </p>
            </div>
        );
    }
}
export default Drawing