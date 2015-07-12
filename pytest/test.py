import webapp2

class Root(webapp2.RequestHandler):
    def get(self):
        self.response.headers['Content-Type'] = 'text/plain'
        self.response.write('Hello, World!\n')

class Test(webapp2.RequestHandler):
    def get(self):
        self.response.write('foo was set to %s' % self.request.get("foo"))

app = webapp2.WSGIApplication([
    ('/', Root),
    ('/test', Test),
], debug=True)
