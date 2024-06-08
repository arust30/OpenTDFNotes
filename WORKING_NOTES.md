# Notes while completing take-home assessment and setting up OpenTDF Platform and cli
 - was tempted to do it on my Arch machine since I prefer thay keyboard, tried for about an hour but wasn't playing nice so ended up just using my MBP
 - installing all dependancies was easy with brew, just wiped this computer so no docker, etc.
 - decided to use Go as a challenge and just to better familiarize myself with the language. I expect this to leave plenty of holes and room for optimizations...and to take 50% longer than it might have otherwise
 - thinking of something simple, maybe only working with plaintext and the local network
 - prompted Mistral's AI a bit for ideas emphasizing _small_ and _proof-of-concept_, also kept appending TDF documentation to have it in its 'attention'
 - landed on a 'Notes' app with encryption from OpenTDF that one could send to other devices on network - stupid and trivial, but something
 - CLI interface with commands like `note new <title>`, `note delete <title>`, etc.
 - Start with outining documentation with what I want, and what I don't want: "Features" and "Anti-Features"
 - Outlined basic directory structure and then started with a Hello World test in Go. Mistral AI made this quite easy, but spent a little time learning about Go just reading the docs and asking Mistral
 - Then read up on the OpenTDF docs for quite some time, trying to grok Platform and the cli tool. Platform necessary for cli tool but most features I won't need
 - docker was being difficult getting setup, the db's worked fine (both the postgres and opentdfdb) but the keycloak daemon was stuck in a restart look, spent a fair bit of time trying to figure out why
 - error messages were mostly Java when reading docker logs, then was getting EOF from OpenTDF,
 - tried playing with the ports (VPN was off and firewall was down)then got connection refused so reverted ports back to defaults
 - tried starting manually, Mistral helped me write:
```
docker run -d --name keycloak \
  -p 8443:8443 -p 8888:8888 \
  -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=changeme \
  -v keycloak_data:/opt/keycloak/data \
  ghcr.io/opentdf/keycloak:sha-8a6d35a
```
 to no avail
 - a few permissions issues (localhost.crv need executable via chmod +x) but didn't resolve issues
 - took a long break. Will try clean install later
--- (spent time away) 
 - started back by reading up on docs for cli tool since that would be much closer to what I am using (assuming I could get the platform up and running)
 - found link for docs in browser but permission denied [here](https://github.com/opentdf/otdfctl/tree/main/docs)
 - so read docs I just used cat: `cd otdfctl/docs/man | cat *`
 - this also gave me fodder to feed prompts to Mistral for better context for `otdfctl <commands>` since it was just guessing at first - still sucks tbh
 - built binary for `otdfctl` manually since installing with go ended with circular imports... probably my fault, but just wanted to get to testing so this was easier
--- (spent time away)
 - wrote some updated specs and usages on the README after thinking a bit more about what to do
 - revised the **very** complicated schema I had before into a single `main.go` file and then a bunch of sub-functions inside. The initial idea might've been better if this were a serious project
 - went through with 'the knife' and just tried to simplify everything as much as possible from where I was conceptually when I started.
 - used Mistral to help simplify a single Go script but did all the README and invocation edits myself - I believe final documentation should be written by human fingers even if LLMs give out the ideas... imho
 - got simple text creation, etc working. Since platform never got up and running, the cli just threw errors so I manually added pretend encryption (reg base64) and did all the testing that way just to isolate the code I wrote 
 - once functions seemed to work, testing minor edge cases, mainly just throwing errors and re-prompting user
 - then went through and refined cruddy comments from before, trying to comment on "why", and not "what", the what is obvious most of the time... as long as you know the language, which in this case I dont (lol)
 - finally prepped simple demo, ran it
 - then purged the local copy, re-cloned it and tested demo again

In conclusion, this took altogether about 7 hours. Most of which was beating my head against setup errors, or re-thinking holistic strategy. Likely only 2 of these hours was spent implementing features, writing code...
