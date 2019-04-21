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
    this.setState({name: event.target.value});
  }
  handlePhoneChange = (event: any) => {
    this.setState({phoneNumber: event.target.value});
  }
  handleCodeChange = (event: any) => {
    this.setState({codes: event.target.value});
  }


  handleSubmit = async (e: any) => {
    e.preventDefault()

    let passwords = ["hi", "hello", "test", "what", "where", "when", "this", "awesome"]
    let hashed = []
    for (const password of passwords) {
      hashed.push(web3.utils.sha3(web3.utils.asciiToHex(password)))
    }
    console.log(hashed)

    try {

      // enable portis
      const accounts = await portis.provider.enable();

      const contractInstance = await this.contract.new("254777777777", "Bob", "0xC28614fEcD3109EFf192DD3cABc7ac9b82C7eD11", passwords, { from: accounts[0] });
    } catch (error) {
      console.error(error);
    }
  };
  render() {
    return (
      <Fragment>

        <div className="container">
          <div className="heading">
            <h3>Welcome to Crypto-USSD!</h3>
            <h4>Create your Brand new Smart Contract by clicking <button onClick={this.openModal}>here!</button></h4>
          </div>
          <Dialog
            open={this.state.open}
            onClose={this.handleClose}
            aria-labelledby="form-dialog-title"
          >
            <DialogTitle id="form-dialog-title">Create Smart Contract</DialogTitle>
            <DialogContent>
              <form onSubmit={this.handleSubmit}>
                <label>
                  Name:
                    <input type="text" name="name"  onChange={this.handleNameChange} value={this.state.name}/>
                </label>
                <label>
                  Your Phone Number:
                    <input type="text" name="phoneNumber" onChange={this.handlePhoneChange} value={this.state.phoneNumber} />
                </label>
                <label>
                  Your codes (comma delimited):
                    <input type="text" name="codes" onChange={this.handleCodeChange} value={this.state.codes} />
                </label>
                <input type="submit" value="Submit" />
              </form>
              <DialogActions>
              </DialogActions>
            </DialogContent>
          </Dialog>
        </div>

      </Fragment >
    );
  }
}
