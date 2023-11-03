
import { create, globSource } from 'ipfs'

let shouldExit = false;

const handleExit = async () => {
  console.log('Gracefully stopping IPFS node...');
  shouldExit = true;
  try {
    await ipfs.stop();
    process.exit(0);
  } catch (err) {
    console.error('Error stopping IPFS:', err);
    process.exit(1);
  }
};

// Register a signal handler to gracefully stop the script
process.on('SIGINT', handleExit);

const startScript = async () => {
  const ipfs = await create()

  for await (const file of ipfs.addAll(globSource('./docs', '**/*'))) {
    console.log(file)
    if (shouldExit) {
      handleExit();
    }
  }

  // All files processed, stop the IPFS node
  await ipfs.stop();
  console.log('IPFS node stopped.');
};

startScript();
