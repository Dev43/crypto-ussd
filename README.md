# crypto-ussd
## Inspiration

My inspiration from this application came from different avenues.
- First and foremost, the main goal that I have is to empower a user to be able to be in charge of their finances. Simplicity achieves that.
- Secondly, I want to accelerate the adoption of crypto and level the playing field as to who can actually use it. Let's face it, right now mostly people with a certain wealth can access a smartphone and sign their transactions on their device, and the others do not. There is about 4.5 billion phones in the world and according to  [this](http://www.lightspeedresearch.com/importance-feature-phones-emerging-markets/): "Of this number, 40% of the global population are smart phones users while the remainder use the feature or ‘dumb’ phones". That's 2.7 billion phones that cannot interact with crypto.

## What it does

The application has 3 parts, a **User** solidity contract that a user needs to deploy at the beginning. I envision this deployment as though it's a "Sign up" to the network. This could be done at points of sale (like MPESA points of sale). Here the user needs to input his phone number and **codes**. By codes I mean hashed secrets, that will be used in a commit-reveal scheme for the "Telco smart contract".

Another part is the Telco smart contract. This smart contract takes care of the logic of transfering tokens between users, and in general doing transactions on behalf of the user. The Telco **cannot** transact the tokens/create CDPs unless the user reveals one of the secrets that they inputted.

The last 2 parts are the USSD Golang Server and the React frontend onboarder application. The USSD server is the one responding to all USSD requests from the cell phone, and creating the transactions as needed/order by the users. 

The react frontend is a small example of the "initial portal" that the onboarding user would see, helping them interact with their contract. The react frontend also uses Portis a possible solution for the user's key management needs, which in certain networks (unfortunately not kovan) uses meta transactions to pay for the user's gas costs.

The application can:
- Send tokens between phone numbers, no need to be registered, you can register in the future and withdraw the coins
- Withdraw tokens to the user's wallet as they wish
- Create a MakerDAO CDP to lock ETH and get DAI tokens, loans in a few clicks!
- Show user's account information such as balance etc
- Deploy a user's User contract from a simple frontend using Portis

## Challenges I ran into

Initially I was going to add Raiden payment channels on it so a User could have multiple payment channels controlled by the Telco. The problem I ran into was that raiden was focused on payment channels between EOA's and not between contracts or even EOA -> contract (which is more general state transition). So I had to abandon that route.
Another problem is that the system of commit-reveal is still not ideal for users, and does not address the possibility of a Telco not executing a transaction on behalf of the user (although the user can retrieve his funds when he wants!)
Also again, this onboarding needed a public/private keypair for the user (stored on Portis) and this still means a lot of complexity for a regular user

## Accomplishments that I'm proud of
It works! It's simple, not too many bells and whisltes and you can have a loan IN 6 CLICKS! That's insane.

## What I learned
USSD applications are actually really cool. When I think that I could send ETH and get a loan with an old style Nokia phone -- there should be more care put into these applications.

ONE SHOULD ALWAYS DO THEIR RESEARCH ON WHAT NETWORKS CERTAIN APPLICATIONS ARE DEPLOYED ON. I changed networks 3 times, from ropsten to goerli to kovan to do this, a huge waste of time.

PS: MakerDao's contract names are terrible, I do not know why they went for such complexity.

## What's next for USSD For All

Look into state channels for contract to contract interactions. 


**To access the USSD menu**, click [here](https://simulator.africastalking.com:1517/simulator/ussd), input the phone number 254777777777, choose USSD and then the channel number *384*7007#