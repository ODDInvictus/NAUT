import type { LayoutServerLoad } from './$types';

export const load = (async () => {
  return {
    version: {
      number: '1.0.0',
      link: 'https://github.com/ODDInvictus/NAUT'
    }
  };
}) satisfies LayoutServerLoad;