## Inspiration
This project initially began as an attempt to implement a chain that could function as a balancer style vault abstraction to facilitate interchain flashloans. However, that was an ambitious goal for a hackathon project even one month away from the deadline. As the deadline approached and several questions about how to implement this functionality were still hanging in the air, we made the decision to approach what seemed to be the most pressing problem, direct access to the data of other chains. Since not all chains implements the packet, or wasm Cosmos SDK modules we had to figure out some way to get access to another chains data. 

## What it does
In its current form it queries the osmosis chain for pool data directly through access to an osmosis node, so that data can be used directly by the chain. Essentially, a separate Cosmos SDK chain can act as though it has access to another chains CLI in order to query data quickly.

## How we built it
We used starport to quickly spin up another chain, defined a message which would allow for parsing arbitrary data and then wrote a keeper function as a test that the data was being passed correctly.

## Challenges we ran into
We are still wondering if this is a valid Cosmos SDK pattern and whether this introduces unnecessary security risk.

## Accomplishments that we're proud of
We are proud of finding a way to communicate between chains that was not very well documented.

## What we learned
We learned a way to force interchain communication where there seemed to be no other options. This has has led us to ask more questions though. Should we continue using this method? Does this align with the IBC protocol?

## What's next for blakyub
- Research how to implement this functionality using an agoric contract, as opposed to starting up a new chain (using existing infrastructure would make onboarding the project more viable)
- Add functionality for broadcasting transactions( in order to be able to perform actions such as swaps & liquidations on behalf of the flashloan, this should be possible by creating a transaction programmatically and broadcasting it using grpcurl)
- engage with the open source developer community in order to build out the full product
