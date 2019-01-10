# Explaining Blockchain and Tendermint to a Developer - Blog Post

I recently had a conversation with my friend during which the subject of blockchain and [Tendermint](https://tendermint.com/) came up. Since I already started exploring both of them, I was able to answer some of the questions he had. Most of us are familiar with the term "blockchain" because of the hype and popularity Bitcoin gained, and my friend was no different. In this blog post, I'll list the questions he asked and give my answers so that I can help all of you understand these topics better.

## What Was the Prime Problem that Bitcoin Solved?
It solved the atomic broadcast problem in a public adversarial setting through clever use of economics.

## What is Atomic Broadcast?
It is a message delivery paradigm which has the following properties:

- Validity: If a correct process broadcasts message m, it eventually delivers m
- Agreement: if a correct process delivers message m all correct processes eventually deliver m
- Integrity: message m is only delivered once only if broadcast by its sender
- Total order: if correct process p and q deliver messages m and m', then p delivers m before m' if q delivers m before m'

## Why Are There Blocks in Blockchain ?
Consensus Algorithms typically commit transactions one at a time by design and implement batching after it. This provides two optimizations which in turn give more throughput and fault tolerance.
- Bandwidth Optimization: Every commit requires two rounds of communication across all validators, batching transactions in blocks amortizes the cost of a commit over all transactions in the block.
- Integrity Optimization: The hash chain of blocks forms an immutable data structure much like a git repository. Enabling authenticity checks for sub-states at any point in history.

## What is Tendermint?
Tendermint is a secure state machine replication in blockchain paradigm. It makes Byzantine-fault Tolerant (BFT) - Atomic Broadcast more accountable and if safety is violated it's always possible to find out who acted maliciously.

## What is a Replicated State Machine?
Many deterministic [state machines](https://www.smashingmagazine.com/2018/01/rise-state-machines/) are replicated across different processes but It acts as a single state machine even if some of these processes fail.

State machines accept inputs called transactions and according to the validity of a transaction, a state transition can occur or not. The transaction is an atomic operation and the state transition logic is governed by the state transition function (application logic).

## What is a Byzantine-fault Tolerant (BFT) Algorithm?
Traditionally consensus protocols which are able to tolerate malicious behaviour are called Byzantine Fault Tolerant Algorithms. 
- Crash faults are easier to handle, as no process can lie to another process. They can typically handle half of the system failing. Systems which only tolerate crash faults can operate via simple majority rule, and therefore typically tolerate simultaneous failure of up to half of the system. If the number of failures the system can tolerate is f, such systems must have at least 2f + 1 processes.
- Byzantine failures are more complicated. In a system of 2f +1 processes,if f are Byzantine, they can co-ordinate to say arbitrary things to the otherf + 1 processes. For instance, suppose we are trying to agree on the value of a single bit, and f = 1, so we have N = 3 processes, A, B, and C, where C is Byzantine. C can tell A that the value is 0 and tell B that it's 1. If A agrees that its 0, and B agrees that its 1, then they will both think they have a majority and commit, thereby violating the safety condition. Hence, the upper bound on faults tolerated by a Byzantine system is strictly lower than a non-Byzantine one.

## How Does Tendermint Relate to Paxos and Raft?
Tendermint is a BFT Algorithm that is it can handle processes behaving arbitrarily and can survive 1/3 of the machines becoming malicious.

Paxos and Raft are not BFT Algorithms so they can only tolerate processes failing and can tolerate up to 1/2 of the processes failing.

## What is ABCI?
Tendermint follows a modular architecture. It abstracts away the complexity behind creating the "blockchain". It makes it easy to build decentralized applications on top of tendermint which is very easy for developers.

Developers only worry about the application logic and nothing else. The application and tendermint-core communicate through an interface called the Application Blockchain Communication Interface(ABCI). To be a valid ABCI application the following interface has to be implemented into it: 

```go
type Application interface {
	// Info/Query Connection
	Info(RequestInfo) ResponseInfo                // Return application info
	SetOption(RequestSetOption) ResponseSetOption // Set application option
	Query(RequestQuery) ResponseQuery             // Query for state

	// Mempool Connection
	CheckTx(tx []byte) ResponseCheckTx // Validate a tx for the mempool

	// Consensus Connection
	InitChain(RequestInitChain) ResponseInitChain    // Initialize blockchain with validators and other info from TendermintCore
	BeginBlock(RequestBeginBlock) ResponseBeginBlock // Signals the beginning of a block
	DeliverTx(tx []byte) ResponseDeliverTx           // Deliver a tx for full processing
	EndBlock(RequestEndBlock) ResponseEndBlock       // Signals the end of a block, returns changes to the validator set
	Commit() ResponseCommit                          // Commit the state and return the application Merkle root hash
}
```
Example ABCI application implementations can be found [on Tendermit GitHub page](https://github.com/tendermint/tendermint/blob/master/abci/example/kvstore/kvstore.go#L59) and [on Hashnode's Mint project](https://github.com/hashnode/mint).
## Where Can I Find More Information About This ?
The [Tendermint docs](https://tendermint.com/docs/) are a very good resource and the [thesis paper](https://allquantor.at/blockchainbib/pdf/buchman2016tendermint.pdf) by Ethan Buchman(co-founder of Tendermint) is a good reference which explains the different protocols in detail.

## Conclusion
This conversation lasted for about half an hour and my friend was convinced enough to try out Tendermint. He was so impressed by the abstraction it provides for building decentralized applications. He has no intentions to go into the theoretical complexity of blockchain and just wants to make an app. I'm pretty excited and looking forward to all the cool apps he is going to build. Let me know if you have any questions about this topic in the comment section below 🙂
