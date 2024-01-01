import { PrismaClient } from '@prisma/client';

const prisma = new PrismaClient()

async function main() {
  await prisma.version.upsert({
    where: { id: 1 },
    update: {},
    create: {
      id: 1,
      number: '0.0.1-dev',
      link: '/',
      releaseLink: '/',
      buildDate: new Date()
    }
  })

  const service = await prisma.service.upsert({
    where: {
      slug: 'ibs3'
    },
    update: {},
    create: {
      slug: 'ibs3',
      name: 'ibs3',
      type: 'http',
      url: 'http://localhost:5173/health',
      status: true,
      lastSeen: new Date()
    }
  })

  for (let i = 1; i <= 50; i++) {
    // Get the day 50 - i days ago.
    const createdAt = new Date()
    createdAt.setDate(createdAt.getDate() - i);

    let status = true

    // some random downtime
    if (i > 20 && i < 30 || i === 44) {
      status = false
    }

    await prisma.uptime.upsert({
      where: {
        id: i
      },
      update: {},
      create: {
        id: i,
        serviceSlug: service.slug,
        createdAt,
        status,
      }
    })
  }

}
main()
  .then(async () => {
    await prisma.$disconnect()
  })
  .catch(async (e) => {
    console.error(e)
    await prisma.$disconnect()
    process.exit(1)
  })
