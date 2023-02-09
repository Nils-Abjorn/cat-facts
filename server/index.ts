import fastify from 'fastify'
import path from 'path'

const server = fastify()

// server.register(require('@fastify/static'), {})

server.get('/ping', async (request, reply) => {
  return 'pong\n'
})

server.listen({ port: 8080 }, (err, address) => {
  if (err) {
    console.error(err)
    process.exit(1)
  }
  console.log(__dirname)
  console.log(`Server listening at ${address}`)
})
