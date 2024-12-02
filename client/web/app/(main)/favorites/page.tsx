import React from 'react';
import { currentUser } from '@clerk/nextjs/server';

const Favorites = async () => {
  const user = await currentUser();

  return (
    <div className="text-2xl text-center font-bold text-primary mt-4">
      Hello {user?.firstName} {user?.lastName}!
    </div>
  );
};

export default Favorites;
