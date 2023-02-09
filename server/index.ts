import fastify from "fastify"
import fastifyStatic from "@fastify/static"

const server = fastify()

server.register(fastifyStatic, {
  root: __dirname,
})

server.get("/", async (request, reply) => {
  return reply.sendFile("index.html") // serving path.join(__dirname, 'public', 'myHtml.html') directly
})

server.get("/ping", async (request, reply) => {
  return "pong\n"
})

server.listen({ port: 8080 }, (err, address) => {
  if (err) {
    console.error(err)
    process.exit(1)
  }
  console.log(__dirname)
  console.log(`Server listening at ${address}`)
})
