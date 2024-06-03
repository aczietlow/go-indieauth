High level goals over the next couple of weeks

[ ] Scaffolding
[ ] Respond with oauth metadata
[ ] Buildout authorization endpoint
    [ ] MUST use PKCE [RFC7636](https://www.rfc-editor.org/rfc/rfc7636)
    [ ] Build code_challenge component 
        - will consider backwards compatability in future by not requiring this
    [ ] Implement Client Discovery using the client id (URL)
    [ ] Validation of authorization request
        - client provided a `me` query parameter
            - either exact user enter value or the value of URL canonicalization
        - able to fetch client information via client discovery
        - client_id URL scheme, host, or port should either match the values of redirect_uri or verify the redirect links of the client 
            - `<link rel="redirect_uri" href="/redirect">