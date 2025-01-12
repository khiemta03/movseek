import axiosInstanceLLM from '@/lib/axios.config.llm';
import { TMDB_API } from '@/utils/constants';

export const fetchLLMRetriever = (collection_name: string, query: string, amount: number, threshold: number) => {
  return axiosInstanceLLM.get(TMDB_API.LLM_RETRIEVER(collection_name, query, amount, threshold));
};
