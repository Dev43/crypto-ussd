import React, { Fragment } from 'react';
import 'react-toastify/dist/ReactToastify.css';
import { web3, portis, web3networks } from '../services/web3';
import User from "../../User.json"
const contract = require('truffle-contract');
import Button from "@material-ui/core/Button";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogContentText from "@material-ui/core/DialogContentText";
import DialogTitle from "@material-ui/core/DialogTitle";

export default class CryptoUSSD extends React.Component<{}, any> {
  contract: any
  public state = {
    open: false,
    phoneNumber: '',
    name: '',
    codes: '',
    address: ''
  };
  async componentDidMount() {
    this.contract = contract(User);
    this.contract.setProvider(web3.currentProvider);
  }

  public openModal = () => {
    this.setState({ open: true });
  }

  public handleClose = () => {
    this.setState({ open: false });
  }

  handleNameChange = (event: any) => {
    this.setState({ name: event.target.value });
  }
  handlePhoneChange = (event: any) => {
    this.setState({ phoneNumber: event.target.value });
  }
  handleCodeChange = (event: any) => {
    this.setState({ codes: event.target.value });
  }


  handleSubmit = async (e: any) => {
    e.preventDefault()

    // let passwords = ["hi", "hello", "test", "what", "where", "when", "this", "awesome"]
    let passwords = this.state.codes.split(",")
    let hashed = []
    for (const password of passwords) {
      hashed.push(web3.utils.sha3(web3.utils.asciiToHex(password)))
    }
    console.log(hashed)

    try {

      // enable portis
      const accounts = await portis.provider.enable();

      const tx = await this.contract.new(this.state.phoneNumber, this.state.name, "0xC28614fEcD3109EFf192DD3cABc7ac9b82C7eD11", passwords, { from: accounts[0] });
      console.log(tx)
      this.setState({address: tx.address})
      // const contractInstance = await this.contract.new("254777777777", "Bob", "0xC28614fEcD3109EFf192DD3cABc7ac9b82C7eD11", passwords, { from: accounts[0] });
    } catch (error) {
      console.error(error);
    }
  };
  render() {
    return (
      <Fragment>

        <div className="container">
          <div className="heading">
            <h1>Welcome to Crypto-USSD!</h1>
            <h4>Before we begin, we need to get you into the amazing decentralized system. Please click <button onClick={this.openModal}>here!</button></h4>
          </div>

        {this.state.address && <div>Wooo you contract is deployed at {this.state.address}</div>}

          <Dialog
            open={this.state.open}
            onClose={this.handleClose}
            aria-labelledby="form-dialog-title"
          >
            <DialogTitle id="form-dialog-title">You're one step away from the decentralized world!</DialogTitle>
            <DialogContent>
              <p>
                For you to be able to enjoy the benefits of the decentralized world, we need you to create a smart contract. Don't worry,
                you won't have to do it yourself, we are helping you along the way.
              </p>
              <DialogActions>
                <form onSubmit={this.handleSubmit}>
                  <label> 
                    Name:
                    <input type="text" name="name" onChange={this.handleNameChange} value={this.state.name} />
                  </label>
                  <br />
                  <br />
                  Now input your phone number, kindly ommit the '+' sign at the beginning
                  <br />
                <label>
                  Phone Number:
                    <input type="text" name="phoneNumber" onChange={this.handlePhoneChange} value={this.state.phoneNumber} />
                  </label>
                  <br />
                  <br />
                  Now we need you to input "codes" in here. These codes are <b>VERY</b> important to remember. They are how you, and only you will be able to 
                  interact with the blockchain (these codes will need to be replenished after they are all used up). Input them like so "hi,hello..." 
                  but make sure you codes are long and no one can guess them!
                  <br />
                <label>
                    Your codes (comma delimited):
                    <input type="text" name="codes" onChange={this.handleCodeChange} value={this.state.codes} />
                  </label>
                  <br />
                  <input type="submit" value="Submit" />
                </form>
              </DialogActions>
            </DialogContent>
          </Dialog>
        </div>

      </Fragment >
    );
  }
}
