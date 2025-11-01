#!/usr/bin/env node

// NOTE: NGINX reverse proxy

//[ IMPs ]/////////////////////////////////////////////////////////////////////
// IMPs - ExtLib
import Fastify from 'fastify'
import helmet from '@fastify/helmet';
// IMPs - local
// import apiRoute from './routes/api.mjs';
import {
  fastOpts,
  helmetOpts,
  corsOpts,
  baseHttpOpts,
  page404,
} from '../serverData.mjs';
import applySandhi from './main.mjs';


//[ DATA ]/////////////////////////////////////////////////////////////////////
const httpOpts = { ...baseHttpOpts, port: 9000 };


//[ INIT ]/////////////////////////////////////////////////////////////////////
const fastify = Fastify(fastOpts);
const notFoundHandler = (_, res) => res.code(404).type('text/html').send(page404);

fastify.setNotFoundHandler(notFoundHandler);


//[ FUNC ]/////////////////////////////////////////////////////////////////////
const errHandler = err => {
  if (err) {
    fastify.log.error(err);
    process.exit(1);
  }
};

fastify.post('/api/transform', async (req, res) => {
  console.log('1. Full request.body:', req.body)
  console.log('2. Type of body:', typeof req.body)

  const { inputTxt } = req.body;
  console.log('3. Extracted inputTxt:', inputTxt)
  console.log('4. Type of inputTxt:', typeof inputTxt)

  if (!inputTxt)
    return res.code(400).send({ error: 'inputTxt is required' });

  const outputTxt = applySandhi(inputTxt);

  res.send(outputTxt);
});


//[ MAIN ]//////////////////////////////////////////////////////////////////////
fastify.register(helmet, helmetOpts);
// fastify.register(apiRoute, { prefix: '/api' });


//[ MAIN-INIT ]/////////////////////////////////////////////////////////////////
fastify.listen(httpOpts, errHandler);
