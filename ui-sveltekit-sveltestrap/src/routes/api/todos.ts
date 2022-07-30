/** @type {import('@sveltejs/kit').RequestHandler} */
export async function GET() {
  return {
    status: 200,
    headers: {
      'access-control-allow-origin': '*'
    },
    body: {
      number: Math.random()
    }
  };
}

export async function POST({ request }) {
  return {
    status: 201,
    body: request.body
  }
}

export async function DELETE({ request }) {
  return {
    status: 201,
    body: request.body
  }
}

