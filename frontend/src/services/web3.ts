import Portis from "@portis/web3";
import Web3 from "web3";

export const portis = new Portis(
  "2afe15d5-e9af-4fba-8d70-68f46c1343a1",
  "kovan",
  { gasRelay: false }
);
export const web3 = new Web3(portis.provider);

export const web3networks: {
  [key: number]: { name: string; explorer: string };
} = {
  1: { name: "Ethereum", explorer: "https://etherscan.io" },
  3: { name: "Ethereum (ropsten)", explorer: "https://ropsten.etherscan.io" },
  4: {
    name: "Ethereum (rinkeby)",
    explorer: "https://blockscout.com/eth/rinkeby"
  },
  5: {
    name: "Ethereum (goerli)",
    explorer: "https://blockscout.com/eth/goerli"
  },
  42: {
    name: "Ethereum (kovan)",
    explorer: "https://blockscout.com/eth/kovan"
  },
  61: {
    name: "Ethereum Classic",
    explorer: "https://blockscout.com/etc/mainnet"
  },
  77: { name: "POA (sokol)", explorer: "https://blockscout.com/poa/sokol" },
  99: { name: "POA", explorer: "https://blockscout.com/poa/core" },
  100: { name: "XDAI", explorer: "https://blockscout.com/poa/dai" }
};
